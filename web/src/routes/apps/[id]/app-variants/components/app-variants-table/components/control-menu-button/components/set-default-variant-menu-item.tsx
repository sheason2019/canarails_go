import { components } from '@/api/api-gen';
import useApi from '@/common/use-api';
import useApp from '@/routes/apps/[id]/hooks/use-app';
import { ListItemText, MenuItem } from '@mui/material';
import { useSnackbar } from 'notistack';
import useSWRMutation from 'swr/mutation';

interface Props {
  appVar: components['schemas']['AppVariant'];
  onClose(): void;
}

export default function SetDefaultVariantMenuItem({ appVar, onClose }: Props) {
  const { data, mutate } = useApp();
  const app = data?.data;
  const api = useApi();
  const { enqueueSnackbar } = useSnackbar();

  const { trigger } = useSWRMutation(
    ['put-app', app?.id],
    async () => {
      if (!app) throw new Error('cannot find app');

      await api.PUT('/api/app/{id}', {
        params: {
          path: { id: Number(app?.id) },
        },
        body: {
          ...app,
          defaultVariantId: appVar.id,
        },
      });
    },
    {
      onSuccess() {
        enqueueSnackbar('设置默认流量泳道成功', { variant: 'success' });
        onClose();
        mutate();
      },
    },
  );

  return (
    <MenuItem onClick={() => trigger()}>
      <ListItemText>设置为默认流量泳道</ListItemText>
    </MenuItem>
  );
}
