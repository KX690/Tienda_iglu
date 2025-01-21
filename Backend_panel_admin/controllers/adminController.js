import Admin from '../models/Admin.js';
import Product from '../models/Product.js';
import Order from '../models/Order.js';
import bcrypt from 'bcryptjs';
import jwt from 'jsonwebtoken';


export const renderLogin = (req, res) => {
  res.render('admin/login');
};

export const loginAdmin = async (req, res) => {
  try {
    const { email, password } = req.body;
    
    const admin = await Admin.findOne({ email });
    
    if (!admin) {
      return res.status(401).json({
        error: 'Credenciales inválidas'
      });
    }

    const isValid = await admin.comparePassword(password);
    if (!isValid) {
      return res.status(401).json({
        error: 'Credenciales inválidas'
      });
    }

    const token = jwt.sign(
      { id: admin._id, email: admin.email },
      process.env.JWT_SECRET,
      { expiresIn: '1d' }
    );
    
    return res.status(200).json({
      token,
      redirect: '/admin/products'
    });
    
  } catch (error) {
    return res.status(500).json({
      error: 'Error al iniciar sesión'
    });
  }
};

// Mostrar formulario de registro
export const showRegister = (req, res) => {
  res.render('admin/register');
};

// Crear nuevo admin
export const createAdmin = (req, res) => {
  const { name, email, password } = req.body;
  const newAdmin = new Admin({
    username: name,
    email: email,
    password: bcrypt.hashSync(password, 10) // Encriptar contraseña
  });
  
  newAdmin.save().then(() => res.redirect('/admin/login'));
};

// Listar productos
export const listProducts = async (req, res) => {
  try {
    const products = await Product.find();
    res.render('admin/products', { products });
  } catch (error) {
    res.render('admin/products', { 
      error: 'Error al cargar los productos',
      products: []
    });
  }
};

// Mostrar formulario de crear producto
export const showCreateProduct = (req, res) => {
  res.render('admin/create-product');
};

// Crear producto
export const createProduct = (req, res) => {
  const newProduct = new Product(req.body);
  newProduct.save().then(() => res.redirect('/admin/products'));
};

// Mostrar formulario de editar producto
export const showEditProduct = (req, res) => {
  Product.findById(req.params.id).then(product => {
    if (!product) return res.redirect('/admin/products');
    res.render('admin/create-product', { product });
  });
};

// Actualizar producto
export const updateProduct = (req, res) => {
  Product.findByIdAndUpdate(req.params.id, req.body).then(() => {
    res.redirect('/admin/products');
  });
};

// Eliminar producto
export const deleteProduct = (req, res) => {
  Product.findByIdAndDelete(req.params.id).then(() => {
    res.redirect('/admin/products');
  });
};

// Órdenes
export const listOrders = (req, res) => {
  Order.find().populate('products.product').then(orders => {
    res.render('admin/orders', { orders, token: req.query.token });
  });
};

export const showOrderDetails = (req, res) => {
  Order.findById(req.params.id).populate('products.product').then(order => {
    res.render('admin/order-details', { order, token: req.query.token });
  });
};

export const updateOrderStatus = (req, res) => {
  Order.findByIdAndUpdate(req.params.id, { status: req.body.status }).then(() => {
    res.redirect(`/admin/orders`);
  });
};
