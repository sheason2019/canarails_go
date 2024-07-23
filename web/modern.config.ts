import { appTools, defineConfig } from '@modern-js/app-tools';
import { tailwindcssPlugin } from '@modern-js/plugin-tailwindcss';

// https://modernjs.dev/en/configure/app/usage
export default defineConfig({
  runtime: {
    router: true,
  },
  plugins: [
    appTools({
      bundler: 'experimental-rspack',
    }),
    tailwindcssPlugin(),
  ],
  tools: {
    devServer: {
      proxy: {
        '/api': {
          target: 'http://localhost:3000/',
          changeOrigin: true,
        },
      },
    },
  },
});
