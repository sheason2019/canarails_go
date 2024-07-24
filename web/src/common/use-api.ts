import createClient from 'openapi-fetch';
import { paths } from '@/api/api-gen';
import { useMemo } from 'react';
import { useSnackbar } from 'notistack';

interface Props {
  toastWhenError?: boolean;
}

export default function useApi(props?: Props) {
  const toastWhenError = props?.toastWhenError ?? true;
  const { enqueueSnackbar } = useSnackbar();

  return useMemo(() => {
    const api = createClient<paths>();
    api.use({
      async onResponse({ response }) {
        if (!response.ok) {
          const errMsg = await response.text();
          if (toastWhenError) {
            enqueueSnackbar(`网络请求错误: ${errMsg}`, { variant: 'error' });
          }
          throw new Error(errMsg);
        }
      },
    });

    return api;
  }, [toastWhenError]);
}
