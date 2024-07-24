import {
  Card,
  CardContent,
  CardActions,
  IconButton,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Button,
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';
import useDialog from '@/common/use-dialog';
import useSWRMutation from 'swr/mutation';
import useApp from '@/routes/apps/[id]/hooks/use-app';
import { useSnackbar } from 'notistack';
import useApi from '@/common/use-api';

interface Props {
  hostname: string;
}

export default function AppHostnameCard({ hostname }: Props) {
  const api = useApi();
  const { enqueueSnackbar } = useSnackbar();
  const { isOpen, onClose, onOpen } = useDialog();
  const { data, mutate } = useApp();

  const { trigger } = useSWRMutation(
    ['delete-hostname', hostname],
    async () => {
      const app = data?.data;
      if (!app) throw new Error('cannot get app');

      await api.PUT('/api/app', {
        body: {
          ...app,
          hostnames: app.hostnames.filter(item => item !== hostname),
        },
      });
    },
    {
      onSuccess() {
        enqueueSnackbar('删除成功', { variant: 'success' });
        onClose();
        mutate();
      },
    },
  );

  return (
    <>
      <Card sx={{ display: 'flex' }}>
        <CardContent sx={{ flex: 1 }}>{hostname}</CardContent>
        <CardActions>
          <IconButton onClick={onOpen}>
            <DeleteIcon color="error" />
          </IconButton>
        </CardActions>
      </Card>
      <Dialog open={isOpen} onClose={onClose}>
        <DialogTitle>警告</DialogTitle>
        <DialogContent>
          <DialogContentText>确定要删除域名匹配 {hostname} 吗？</DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={onClose}>取消</Button>
          <Button onClick={() => trigger()} variant="contained" color="error">
            确定
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
}
