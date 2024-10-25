/** @type {import('tailwindcss').Config} */
const colors = require("tailwindcss/colors");

module.exports = {
  content: ["./view/**/*.html"],
  theme: {
    colors: {
      transparent: "transparent",
      current: "currentColor",
      black: colors.black,
      white: colors.white,
      gray: colors.gray,
      cyan: "#79FFF7",
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
  plugins: [],
};
