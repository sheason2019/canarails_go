import createClient from 'openapi-fetch';
import { paths } from './api-gen';

export const api = createClient<paths>();
