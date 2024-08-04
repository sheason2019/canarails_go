import useSWR from 'swr';
import useToken from './use-token';
import useApi from '../use-api';
import { UnloginError } from '../errors/unlogin-error';

export default function useUser() {
  const api = useApi({ toastWhenError: false });
  const { setToken } = useToken();

  const { data, isLoading, mutate } = useSWR(
     'auth/user',
    async () => {
      try {
        return await api.GET('/api/auth');
      } catch (e) {
        if (e instanceof UnloginError) {
          return null;
        }

        throw e;
      }
    },
    {
      shouldRetryOnError: false,
      suspense: true,
    },
  );

  const logout = () => {
    setToken('');
  };

  return {
    user: data?.data,
    isLogined: !!data,
    isLoading,
    logout,
    mutate,
  };
}
