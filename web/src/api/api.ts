import createClient from 'openapi-fetch';
import { paths } from './api-gen';

const api = createClient<paths>();
api.use({
  async onResponse({ response }) {
    if (!response.ok) {
      throw new Error(await response.text());
    }
  },
});

export { api };
