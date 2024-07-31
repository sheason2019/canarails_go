import { components } from '@/api/api-gen';

export type CreateUserTokenRes = components['schemas']['CreateUserTokenRes'];
export type UserToken = Omit<CreateUserTokenRes, 'tokenString'> &
  Partial<Pick<CreateUserTokenRes, 'tokenString'>>;
