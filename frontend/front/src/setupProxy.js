const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = function (app) {
  app.use(
    "/api",
    createProxyMiddleware({
      target: process.env.PROXY_API || "http://localhost:8080",
      changeOrigin: true,
      pathRewrite: {
        "^/api": "/",
      },
    })
  );
};
