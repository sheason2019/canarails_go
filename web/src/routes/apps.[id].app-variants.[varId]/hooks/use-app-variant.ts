import useApi from '@/common/use-api';
import { useParams } from '@modern-js/runtime/router';
import useSWR from 'swr';

export default function useAppVariant() {
  const { varId } = useParams();
  const api = useApi();

  return useSWR(['app-variant', varId], () =>
    api.GET('/api/app-variant/{id}', {
      params: {
        path: {
          id: Number(varId),
        },
      },
    }),
  );
}
