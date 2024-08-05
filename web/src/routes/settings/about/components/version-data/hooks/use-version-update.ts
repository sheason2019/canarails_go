import useApi from '@/common/use-api';
import { useSnackbar } from 'notistack';
import useSWRMutation from 'swr/mutation';

export default function useVersionUpdate() {
  const api = useApi();

  const { enqueueSnackbar } = useSnackbar();

  const { trigger: checkUpdate } = useSWRMutation(
    'check-update',
    () => api.GET('/api/version/update'),
    {
      onSuccess(data) {
        if (data.data?.gitHash) {
          console.log('存在新版');
        } else {
          enqueueSnackbar('当前已是最新版本', { variant: 'success' });
        }
      },
    },
  );

  return {
    checkUpdate,
  };
}
