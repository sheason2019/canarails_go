import { Outlet } from '@modern-js/runtime/router';
import AppHeader from './components/app-header';
import { Container, CssBaseline } from '@mui/material';

import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import './layout.css';
import { RecoilRoot } from 'recoil';
import { SnackbarProvider } from 'notistack';
import UserProvider from '@/common/user/user-provider';

export default function Layout() {
  return (
    <RecoilRoot>
      <CssBaseline />
      <UserProvider>
        <SnackbarProvider>
          <AppHeader />
          <Container>
            <Outlet />
          </Container>
        </SnackbarProvider>
      </UserProvider>
    </RecoilRoot>
  );
}
