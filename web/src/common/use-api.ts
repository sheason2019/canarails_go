import createClient from 'openapi-fetch';
import { paths } from '@/api/api-gen';
import { useMemo } from 'react';
import { useSnackbar } from 'notistack';
import useToken from './user/use-token';
import { UnloginError } from './errors/unlogin-error';

interface Props {
  toastWhenError?: boolean;
}

export default function useApi(props?: Props) {
  const { token } = useToken();

  const toastWhenError = props?.toastWhenError ?? true;
  const { enqueueSnackbar } = useSnackbar();

  return useMemo(() => {
    const api = createClient<paths>();
    api.use({
      onRequest({ request }) {
        request.headers.set('Authorization', token);
      },
      async onResponse({ response }) {
        if (response.status === 401) {
          if (toastWhenError) {
            enqueueSnackbar('当前用户未登录', { variant: 'error' });
          }
          throw new UnloginError('当前用户未登录');
        }

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
  }, [toastWhenError, token]);
}
