package handlers

import (
    "fmt"
   
    "net/http"
    "tienda/database"
    "tienda/models"
    "go.mongodb.org/mongo-driver/bson"
    "context"
    "time"
    "strconv"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleProducts(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var products []models.Product
    cursor, err := database.DB.Collection("products").Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    if err = cursor.All(ctx, &products); err != nil {
        http.Error(w, "Error al decodificar productos", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(w, `
        <h1>Productos</h1>
        <style>
            .product { margin: 20px 0; padding: 10px; border: 1px solid #ddd; }
            .order-form { display: none; margin-top: 10px; }
            input { margin: 5px 0; padding: 5px; }
        </style>
        <script>
            function showOrderForm(productId) {
                document.getElementById('form-' + productId).style.display = 'block';
            }
        </script>
    `)
    
    for _, p := range products {
        fmt.Fprintf(w, `
            <div class="product">
                <div>Nombre: %s - Precio: $%.2f - Stock: %d</div>
                <button onclick="showOrderForm('%s')">Generar Orden</button>
                <form id="form-%s" class="order-form" action="/order" method="POST">
                    <input type="hidden" name="product_id" value="%s">
                    <input type="text" name="name" placeholder="Tu nombre" required><br>
                    <input type="email" name="email" placeholder="Tu email" required><br>
                    <input type="number" name="quantity" value="1" min="1" max="%d" required><br>
                    <input type="submit" value="Confirmar Orden">
                </form>
            </div>
        `, p.Name, p.Price, p.Stock, p.ID.Hex(), p.ID.Hex(), p.ID.Hex(), p.Stock)
    }
}

func HandleOrder(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Obtener datos del formulario
    productID := r.FormValue("product_id")
    quantity, _ := strconv.Atoi(r.FormValue("quantity"))

    // Verificar stock disponible
    objID, err := primitive.ObjectIDFromHex(productID)
    if err != nil {
        http.Error(w, "ID de producto inválido", http.StatusBadRequest)
        return
    }

    var product models.Product
    err = database.DB.Collection("products").FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
    if err != nil {
        http.Error(w, "Producto no encontrado", http.StatusNotFound)
        return
    }

    if product.Stock < quantity {
        http.Error(w, "No hay suficiente stock disponible", http.StatusBadRequest)
        return
    }

    // Actualizar el stock
    update := bson.M{
        "$inc": bson.M{"stock": -quantity},
    }
    _, err = database.DB.Collection("products").UpdateOne(ctx, bson.M{"_id": objID}, update)
    if err != nil {
        http.Error(w, "Error al actualizar el stock", http.StatusInternalServerError)
        return
    }

    // Crear la orden
    orderProduct := models.OrderProduct{
        Product:  objID,
        Quantity: quantity,
        Price:    product.Price,
    }

    order := models.Order{
        ID:        primitive.NewObjectID(),
        Customer: models.Customer{
            Name:  r.FormValue("name"),
            Email: r.FormValue("email"),
        },
        Products:  []models.OrderProduct{orderProduct},
        Total:     product.Price * float64(quantity),
        Status:    models.OrderStatusPending,
        CreatedAt: time.Now(),
    }

    _, err = database.DB.Collection("orders").InsertOne(ctx, order)
    if err != nil {
        // Si falla al crear la orden, revertimos la actualización del stock
        revertUpdate := bson.M{
            "$inc": bson.M{"stock": quantity},
        }
        database.DB.Collection("products").UpdateOne(ctx, bson.M{"_id": objID}, revertUpdate)
        http.Error(w, "Error al guardar la orden", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, `
        <h1>Orden Generada</h1>
        <p>Gracias %s por tu orden!</p>
        <p>Te enviaremos un email a %s con los detalles.</p>
        <p>Cantidad ordenada: %d</p>
        <p>Total a pagar: $%.2f</p>
        <p>Stock restante: %d</p>
        <a href="/products">Volver a Productos</a>
    `, order.Customer.Name, order.Customer.Email, orderProduct.Quantity, order.Total, product.Stock - quantity)
}