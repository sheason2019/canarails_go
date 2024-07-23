import useDialog from '@/common/use-dialog';
import useUser from '@/common/user/use-user';
import {
  Avatar,
  Button,
  CardContent,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  Stack,
  Typography,
} from '@mui/material';

export default function LogoutCard() {
  const { isOpen, onClose, onOpen } = useDialog();
  const { logout } = useUser();

  const handleLogout = () => {
    logout();
    onClose();
  };

  return (
    <>
      <CardContent component={Stack} alignItems="center">
        <Typography
          variant="h6"
          sx={{ textAlign: 'center', fontWeight: 'bold' }}
        >
          用户信息
        </Typography>
        <Avatar sx={{ width: 72, height: 72, mt: 2 }} />
        <Button
          variant="contained"
          color="error"
          sx={{ mt: 3, px: 4 }}
          onClick={onOpen}
        >
          退出登录
        </Button>
      </CardContent>
      <Dialog
        open={isOpen}
        onClose={onClose}
        fullWidth
        maxWidth="xs"
        aria-labelledby="退出登录"
        aria-describedby="确认退出登录"
      >
        <DialogTitle>请确认操作</DialogTitle>
        <DialogContent>
          <DialogContentText>即将退出登录，确定要继续吗？</DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={onClose}>取消</Button>
          <Button onClick={handleLogout} color="error" variant="contained">
            确定
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
}
