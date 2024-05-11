/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
  theme: {
    extend: {
      colors: {
        primary: "#0d9488",
        secondary: "#94390D",
      },
      fontFamily: {
        sans: ["Red Hat Display", "sans-serif"],
      },
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
