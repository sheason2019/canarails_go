import { Alert, Button, Paper, Stack, Typography } from '@mui/material';
import { UserToken } from '../../typings';
import dayjs from 'dayjs';
import Grid from '@mui/material/Unstable_Grid2/Grid2';
import { useSnackbar } from 'notistack';
import UserTokenItemMenuButton from './components/user-token-item-menu-button';

interface Props {
  userToken: UserToken;
}

export default function UserTokenItem({ userToken }: Props) {
  const { enqueueSnackbar } = useSnackbar();

  const handleCopy = async () => {
    if (userToken.tokenString) {
      await navigator.clipboard.writeText(userToken.tokenString);
      enqueueSnackbar('复制成功', { variant: 'success' });
    }
  };

  return (
    <Paper elevation={0} variant="outlined" sx={{ padding: 2 }}>
      <Stack direction="row" alignItems="center">
        <Typography fontWeight="bold" sx={{ flex: 1 }}>
          {userToken.title}
        </Typography>
        <UserTokenItemMenuButton userToken={userToken} />
      </Stack>
      <Typography variant="body1">{userToken.description}</Typography>
      <Grid container>
        <Grid xs={12} md={4}>
          <Typography variant="body2" color="GrayText">
            {dayjs(userToken.expiredAt).format('YYYY-MM-DD HH:mm')} 到期
          </Typography>
        </Grid>
        <Grid xs={12} md={4}>
          <Typography variant="body2" color="GrayText">
            {dayjs(userToken.lastUsedAt).format('YYYY-MM-DD HH:mm')}{' '}
            最后一次使用
          </Typography>
        </Grid>
      </Grid>
      {userToken.tokenString && (
        <Alert
          sx={{ mt: 1, wordBreak: 'break-all' }}
          severity="success"
          action={
            <Button color="inherit" size="small" onClick={handleCopy}>
              COPY
            </Button>
          }
        >
          <div>令牌信息退出页面后将无法再次查看，请妥善保管：</div>
          <div>{userToken.tokenString}</div>
        </Alert>
      )}
    </Paper>
  );
}
