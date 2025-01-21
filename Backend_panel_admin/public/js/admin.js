// Verificar token al cargar la página
document.addEventListener('DOMContentLoaded', () => {
  const token = localStorage.getItem('token');
  const isLoginPage = window.location.pathname === '/admin/login';
  const isRegisterPage = window.location.pathname === '/admin/register';

  if (!token && !isLoginPage && !isRegisterPage) {
    window.location.href = '/admin/login';
    return;
  }
});

// Interceptar todas las peticiones fetch
const originalFetch = window.fetch;
window.fetch = async (...args) => {
  const [url, config = {}] = args;
  const token = localStorage.getItem('token');
  
  // Agregar token a la URL si existe
  const urlObj = new URL(url, window.location.origin);
  if (token) {
    urlObj.searchParams.append('token', token);
  }

  try {
    const response = await originalFetch(urlObj.toString(), config);
    
    if (response.status === 401) {
      localStorage.removeItem('token');
      window.location.href = '/admin/login';
      return;
    }

    return response;
  } catch (error) {
    console.error('Error en fetch:', error);
    throw error;
  }
};
