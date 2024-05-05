/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Red Hat Display", "sans-serif"],
      },
    },
  },
};
