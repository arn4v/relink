const colors = require("tailwindcss/colors");

delete colors.lightBlue;

module.exports = {
  method: "jit",
  purge: ["./src/**/*.{js,jsx,ts,tsx,vue,svelte,html,css}"],
  darkMode: false,
  theme: {
    extend: {
      colors,
      fontFamily: {
        sans: ['"Space Grotesk"', "sans-serif"],
      },
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
