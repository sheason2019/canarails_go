import useApi from '@/common/use-api';
import useSWR from 'swr';

export default function useVersion() {
  const api = useApi();

  return useSWR('app-version', () => api.GET('/api/version'), {
    suspense: true,
  });
}
