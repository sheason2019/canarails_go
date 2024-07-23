import { Box, Card, Container } from '@mui/material';
import LoginCard from './components/login-card';
import useUser from '@/common/user/use-user';
import LogoutCard from './components/logout-card';

export default function Login() {
  const { isLogined } = useUser();

  return (
    <main>
      <Box
        sx={{
          transform: 'translateY(-50%)',
          top: '50%',
          left: 0,
          right: 0,
          position: 'absolute',
          mx: 'auto',
        }}
      >
        <Container maxWidth="xs">
          <Card elevation={2}>
            {isLogined ? <LogoutCard /> : <LoginCard />}
          </Card>
        </Container>
      </Box>
    </main>
  );
}
