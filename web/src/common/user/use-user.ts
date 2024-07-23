import useSWR from 'swr';
import useToken from './use-token';
import { api } from '@/api/api';

export default function useUser() {
  const { token, setToken } = useToken();

  const { data, error, mutate } = useSWR(['auth/user', token], () =>
    api.GET('/api/auth', { params: { header: { authorization: token } } }),
  );

  const logout = () => {
    setToken('');
  };

  return {
    user: data?.data,
    isLogined: !Boolean(error) && !Boolean(data?.error),
    logout,
  };
}
