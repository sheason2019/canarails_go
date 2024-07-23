import { api } from '@/api/api';
import useSWR from 'swr';

export default function useAppList() {
  return useSWR('app-list', () => api.GET('/api/app'));
}
