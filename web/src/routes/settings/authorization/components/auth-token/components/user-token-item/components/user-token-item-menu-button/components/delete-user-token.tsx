import {
  MenuItem,
  ListItemText,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Button,
} from '@mui/material';
import { UserToken } from '../../../../../typings';
import useSWRMutation from 'swr/mutation';
import useApi from '@/common/use-api';
import { useSnackbar } from 'notistack';
import useUserToken from '../../../../../hooks/use-user-token';
import useDialog from '@/common/use-dialog';

interface Props {
  userToken: UserToken;
  onClose(): void;
}

export default function DeleteUserToken({
  userToken,
  onClose: onCloseMenu,
}: Props) {
  const { mutate } = useUserToken();
  const api = useApi();
  const { enqueueSnackbar } = useSnackbar();

  const { trigger } = useSWRMutation(
    ['delete-user-token'],
    () =>
      api.DELETE('/api/user-token/{id}', {
        params: {
          path: { id: userToken.id },
        },
      }),
    {
      onSuccess() {
        enqueueSnackbar('删除成功', { variant: 'success' });
        onClose();
        onCloseMenu();
        mutate();
      },
    },
  );

  const { isOpen, onClose, onOpen } = useDialog();

  const handleClick = () => {
    onOpen();
  };

  return (
    <>
      <MenuItem onClick={handleClick}>
        <ListItemText sx={{ color: 'red' }}>删除令牌</ListItemText>
      </MenuItem>
      <Dialog open={isOpen} onClose={onClose}>
        <DialogTitle>警告</DialogTitle>
        <DialogContent>
          <DialogContentText>
            即将删除权限令牌，删除后的权限令牌无法恢复，确认要删除吗？
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={onClose}>取消</Button>
          <Button onClick={() => trigger()} color="error" variant="contained">
            确认
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
}
