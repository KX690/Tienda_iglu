import mongoose from 'mongoose'

const orderSchema = new mongoose.Schema({
  customer: {
    name: {type: String, required: true},
    email: {type: String,required: true},
    phone: String,
    address: {type: String, required: true},
  },
  products: [{
    product: {type: mongoose.Schema.Types.ObjectId,ref: 'Product',required: true},
    quantity: {type: Number, required: true, min: 1},
    price: Number
  }],
  total: {type: Number, required: true},
  status: {type: String, enum: ['pendiente', 'pagado', 'entregado', 'cancelado']},
  createdAt: {type: Date, default: Date.now }
});

export default mongoose.model('Order', orderSchema);