import useDialog from '@/common/use-dialog';
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
import useSWRMutation from 'swr/mutation';
import useApi from '@/common/use-api';
import { useSnackbar } from 'notistack';
import useAppVariant from '../hooks/use-app-variant';

interface Props {
  header: string;
}

export default function DeleteHeaderMatchButton({ header }: Props) {
  const { isOpen, onClose, onOpen } = useDialog();
  const { data, mutate } = useAppVariant();
  const api = useApi();
  const { enqueueSnackbar } = useSnackbar();

  const { trigger } = useSWRMutation(
    ['app-variant/delete', header],
    async () => {
      const appVar = data?.data;
      if (!appVar) throw new Error('cannot get appVar');

      await api.PUT('/api/app-variant/{id}', {
        params: {
          path: { id: appVar.id },
        },
        body: {
          ...appVar,
          matches: appVar.matches.filter(item => item.header !== header),
        },
      });
    },
    {
      onSuccess() {
        onClose();
        mutate();
        enqueueSnackbar('删除成功', { variant: 'success' });
      },
    },
  );

  return (
    <>
      <IconButton onClick={onOpen} size="small">
        <DeleteIcon color="error" />
      </IconButton>
      <Dialog open={isOpen} onClose={onClose}>
        <DialogTitle>警告</DialogTitle>
        <DialogContent>
          <DialogContentText>
            确定要删除标头匹配 {header} 吗？
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
