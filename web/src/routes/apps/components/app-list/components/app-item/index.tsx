import { components } from '@/api/api-gen';
import {
  Button,
  Card,
  CardContent,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  IconButton,
  Stack,
  Typography,
} from '@mui/material';
import useSWRMutation from 'swr/mutation';
import DeleteIcon from '@mui/icons-material/Delete';
import useDialog from '@/common/use-dialog';
import { api } from '@/api/api';
import useToken from '@/common/user/use-token';
import useAppList from '../../hooks/use-app-list';

interface Props {
  app: components['schemas']['App'];
}

export default function AppItem({ app }: Props) {
  const { token } = useToken();
  const { mutate } = useAppList();
  const { isOpen, onClose, onOpen } = useDialog();

  const { trigger, isMutating } = useSWRMutation(
    ['delete-app', app.id],
    () =>
      api.DELETE('/api/app', {
        params: {
          header: {
            authorization: token,
          },
        },
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
      <Card>
        <CardContent>
          <Typography variant="h5">{app.title}</Typography>
          <Typography variant="body1">简介：{app.description}</Typography>
          <Stack direction="row" justifyContent="space-between" sx={{ mt: 1 }}>
            <IconButton color="error" onClick={onOpen}>
              <DeleteIcon />
            </IconButton>
            <Button variant="contained">查看详情</Button>
          </Stack>
        </CardContent>
      </Card>
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
