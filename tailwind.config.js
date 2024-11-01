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
      typography: (theme) => ({
        DEFAULT: {
          css: {
            code: {
              borderRadius: theme("borderRadius.sm"),
              backgroundColor: theme("colors.gray.300"),
              color: theme("colors.gray.900"),
              fontWeight: "normal",
              padding: theme("spacing.1"),
            },
            "code::before": {
              content: "",
            },
            "code::after": {
              content: "",
            },
            blockQuote: {
              border: "solid",
              borderWidth: "2px",
              borderStyle: "dashed",
              borderRadius: theme("borderRadius.lg"),
              borderColor: theme("colors.yellow"),
              fontStyle: "normal",
            },
            "blockquote p:first-of-type::before": {
              content: "",
            },
            "blockquote p:last-of-type::after": {
              content: "",
            },
          },
        },
      }),
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
