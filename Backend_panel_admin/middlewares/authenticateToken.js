import jwt from 'jsonwebtoken';
import dotenv from 'dotenv';
dotenv.config();

export const authenticateToken = (req, res, next) => {
  // Intentar obtener el token del header o query parameter
  const authHeader = req.headers.authorization;
  const queryToken = req.query.token;
  
  const token = authHeader ? authHeader.split(' ')[1] : queryToken;

  if (!token) {
    // Si la petición espera JSON, enviar error JSON
    if (req.xhr || req.headers.accept?.includes('json')) {
      return res.status(401).json({ error: 'Acceso denegado' });
    }
    // Si no, redirigir al login
    return res.redirect('/admin/login');
  }

  try {
    const verified = jwt.verify(token, process.env.JWT_SECRET);
    req.user = verified;
    next();
  } catch (error) {
    if (req.xhr || req.headers.accept?.includes('json')) {
      return res.status(401).json({ error: 'Token inválido' });
    }
    return res.redirect('/admin/login');
  }
};
