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
import useSWRMutation from 'swr/mutation';
import DeleteIcon from '@mui/icons-material/Delete';
import useDialog from '@/common/use-dialog';
import useAppList from '../../hooks/use-app-list';
import useApi from '@/common/use-api';

interface Props {
  app: components['schemas']['App'];
}

export default function DeleteAppButton({ app }: Props) {
  const api = useApi();
  const { mutate } = useAppList();
  const { isOpen, onClose, onOpen } = useDialog();

  const { trigger, isMutating } = useSWRMutation(
    ['delete-app', app.id],
    () =>
      api.DELETE('/api/app', {
        body: { id: app.id },
      }),
    {
      onSuccess() {
        onClose();
        mutate();
      },
    },
  );

  const handleClose = () => {
    if (!isMutating) onClose();
  };

  return (
    <>
      <IconButton color="error" onClick={onOpen}>
        <DeleteIcon />
      </IconButton>
      <Dialog open={isOpen} onClose={handleClose}>
        <DialogTitle>警告</DialogTitle>
        <DialogContent>
          <DialogContentText>确定要删除 App {app.title} 吗？</DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} disabled={isMutating}>
            取消
          </Button>
          <Button
            variant="contained"
            color="error"
            onClick={() => trigger()}
            disabled={isMutating}
          >
            确定
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
}
