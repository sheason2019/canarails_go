import { Box, Card } from '@mui/material';
import LoginCard from './components/login-card';
import useUser from '@/common/user/use-user';
import LogoutCard from './components/logout-card';

export default function Login() {
  const { isLogined } = useUser();

  return (
    <main>
      <Box
        className="max-w-md px-4 absolute left-0 right-0 mx-auto"
        sx={{ transform: 'translateY(-50%)', top: '50%' }}
      >
        <Card elevation={2}>{isLogined ? <LogoutCard /> : <LoginCard />}</Card>
      </Box>
    </main>
  );
}
