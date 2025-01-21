import express from 'express';
import * as adminController from '../controllers/adminController.js';
import { authenticateToken } from '../middlewares/authenticateToken.js';


const router = express.Router();

// Middlewares para procesar el body
router.use(express.json());
router.use(express.urlencoded({ extended: true }));

// Proteger todas las rutas excepto login y register
router.use((req, res, next) => {
    if (req.path === '/login' || req.path === '/register') {
        return next();
    }
    return authenticateToken(req, res, next);
});
  
// Rutas de autenticación
router.get('/login', adminController.renderLogin);
router.post('/login', adminController.loginAdmin);

// Rutas de registro 
router.get('/register', adminController.showRegister);
router.post('/register', adminController.createAdmin);


// Rutas de productos
router.get('/products', adminController.listProducts);
router.get('/products/create', adminController.showCreateProduct);
router.post('/products/create', adminController.createProduct);
router.get('/products/edit/:id', adminController.showEditProduct);
router.post('/products/edit/:id', adminController.updateProduct);
router.post('/products/delete/:id', adminController.deleteProduct);

// Rutas de órdenes
router.get('/orders', adminController.listOrders);
router.get('/orders/:id', adminController.showOrderDetails);
router.post('/orders/:id/status', adminController.updateOrderStatus);




export default router;
