// Hapus baris ini: const { defineConfig } = require("@vue/cli-service");

module.exports = {
  devServer: {
    port: 12345, // Coba port yang berbeda
    host: "localhost",
  },

  chainWebpack: (config) => {
    config.plugin("define").tap((args) => {
      args[0]._VUE_PROD_HYDRATION_MISMATCH_DETAILS_ = JSON.stringify(false);
      args[0]._VUE_OPTIONS_API_ = JSON.stringify(true);
      args[0]._VUE_PROD_DEVTOOLS_ = JSON.stringify(false);
      return args;
    });
  },
};
