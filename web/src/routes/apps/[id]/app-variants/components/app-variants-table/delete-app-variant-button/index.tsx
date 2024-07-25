import { components } from '@/api/api-gen';
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  IconButton,
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';
import useDialog from '@/common/use-dialog';
import useSWRMutation from 'swr/mutation';
import useApi from '@/common/use-api';
import { useSnackbar } from 'notistack';
import useAppVariants from '../hooks/use-app-variants';

interface Props {
  appVar: components['schemas']['AppVariant'];
}

export default function DeleteAppVariantButton({ appVar }: Props) {
  const { mutate } = useAppVariants();
  const { isOpen, onOpen, onClose } = useDialog();
  const api = useApi();
  const { enqueueSnackbar } = useSnackbar();

  const { trigger } = useSWRMutation(
    ['delete-app-variant', appVar.id],
    () =>
      api.DELETE('/api/app-variant/{id}', {
        params: {
          path: {
            id: appVar.id,
          },
        },
      }),
    {
      onSuccess() {
        enqueueSnackbar('删除流量泳道成功', { variant: 'success' });
        onClose();
        mutate();
      },
    },
  );

  return (
    <>
      <IconButton onClick={onOpen}>
        <DeleteIcon color="error" />
      </IconButton>
      <Dialog open={isOpen} onClose={onClose}>
        <DialogTitle>警告</DialogTitle>
        <DialogContent>
          <DialogContentText>
            确定要删除流量泳道 {appVar.title} 吗？
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={onClose}>取消</Button>
          <Button variant="contained" color="error" onClick={() => trigger()}>
            确定
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
}
