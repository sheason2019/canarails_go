import useSWR from 'swr';
import useToken from './use-token';
import useApi from '../use-api';

export default function useUser() {
  const api = useApi({ toastWhenError: false });
  const { token, setToken } = useToken();

  const { data, error, isLoading } = useSWR(
    ['auth/user', token],
    () => api.GET('/api/auth'),
    {
      shouldRetryOnError: false,
    },
  );

  const logout = () => {
    setToken('');
  };

  return {
    user: data?.data,
    isLogined: !Boolean(error) && !Boolean(data?.error) && !isLoading,
    logout,
  };
}
