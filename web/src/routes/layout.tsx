import { Outlet } from '@modern-js/runtime/router';
import AppHeader from './components/app-header';
import { ScopedCssBaseline } from '@mui/material';

import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import './layout.css';
import { RecoilRoot } from 'recoil';
import { SnackbarProvider } from 'notistack';

export default function Layout() {
  return (
    <RecoilRoot>
      <SnackbarProvider>
        <ScopedCssBaseline>
          <AppHeader />
          <Outlet />
        </ScopedCssBaseline>
      </SnackbarProvider>
    </RecoilRoot>
  );
}
