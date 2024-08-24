/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.templ", "./**/*.go"],
  darkMode: "selector",
  theme: {
    extend: {
      fontFamily: {
        sans: ["Lato", "sans-serif"],
      },
    },
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        light: {
          ...require("daisyui/src/theming/themes")["light"],
          primary: "#0a776e",
          secondary: "#0369a1",
          accent: "#d97706",
          warning: "#f7b708",
        },
        dark: {
          ...require("daisyui/src/theming/themes")["dark"],
          primary: "#0a776e",
          secondary: "#0369a1",
          accent: "#b35c00",
          warning: "#f7b708",
          "base-content": "#ECECEC",
        },
      },
    ],
  },
};
