import useApi from '@/common/use-api';
import useSWR from 'swr';

export default function useUserToken() {
  const api = useApi();

  return useSWR(['user-token'], () => api.GET('/api/user-token'));
}
