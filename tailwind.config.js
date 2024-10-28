/** @type {import('tailwindcss').Config} */
const colors = require("tailwindcss/colors");

module.exports = {
  content: ["./view/**/*.html", "./public/**/*.js", "./content/**/*.md"],
  theme: {
    colors: {
      transparent: "transparent",
      current: "currentColor",
      black: colors.black,
      white: colors.white,
      gray: colors.gray,
      cyan: {
        DEFAULT: "#79FFF7",
        dark: "#6EE6DF",
      },
      pink: "#FF66A9",
      yellow: "#FDD600",
      dark: "#0F0F0F",
      light: "#F9F9F9",
    },
    extend: {
      boxShadow: {
        solid: "6px 6px 0 rgba(0, 0, 0, 0.25)",
      },
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
