import useApi from '@/common/use-api';
import useSWR from 'swr';

export default function useAppList() {
  const api = useApi();
  return useSWR('app-list', () => api.GET('/api/app'), {suspense: true});
}
