/** @type {import('tailwindcss').Config} */

const config = {
  theme: {
    extend: {
      typography: (theme) => ({
        DEFAULT: {
          css: {
            code: {
              borderRadius: theme("borderRadius.sm"),
              backgroundColor: theme("colors.gray.300"),
              color: theme("colors.gray.900"),
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

export default config;
