import { Paper, Stack, Typography } from '@mui/material';
import useUserToken from './hooks/use-user-token';
import AddUserTokenButton from './components/add-user-token-button';
import { useMemo, useState } from 'react';
import { CreateUserTokenRes, UserToken } from './typings';
import UserTokenItem from './components/user-token-item';

export default function AuthToken() {
  const { data } = useUserToken();

  const [newToken, setNewToken] = useState<CreateUserTokenRes>();

  const userTokens: UserToken[] = useMemo(() => {
    const userTokens: UserToken[] = (data?.data ?? []).filter(
      item => item.id !== newToken?.id,
    );
    if (newToken) {
      userTokens.unshift(newToken);
    }
    return userTokens;
  }, [data?.data, newToken]);

  return (
    <>
      <Stack direction="row" justifyContent="space-between" sx={{ my: 2 }}>
        <Typography variant="h6">权限令牌</Typography>
        <AddUserTokenButton onCreate={data => setNewToken(data)} />
      </Stack>
      {!userTokens?.length ? (
        <Stack sx={{ height: 48 }} justifyContent="center" alignItems="center">
          <Typography variant="body2" color="GrayText">
            暂无数据
          </Typography>
        </Stack>
      ) : (
        <Stack spacing={2}>
          {userTokens.map(ut => (
            <UserTokenItem key={ut.id} userToken={ut} />
          ))}
        </Stack>
      )}
    </>
  );
}
