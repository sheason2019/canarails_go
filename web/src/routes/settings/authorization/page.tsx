import { Typography } from '@mui/material';
import AuthToken from './components/auth-token';

export default function AuthorizationSetting() {
  return (
    <main>
      <Typography sx={{ mb: 3 }} variant="h5">
        权限
      </Typography>
      <AuthToken />
    </main>
  );
}
