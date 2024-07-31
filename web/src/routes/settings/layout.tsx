import { Typography, List, Box } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2/Grid2';
import { Outlet, useLocation } from '@modern-js/runtime/router';
import NavigationItem from '../components/app-header/components/navigation-item';
import KeyIcon from '@mui/icons-material/Key';
import PersonIcon from '@mui/icons-material/Person';
import { useMemo } from 'react';

enum SettingRoutes {
  profile,
  authorization,
}

export default function Layout() {
  const location = useLocation();
  const currentRoute = useMemo(() => {
    const current = location.pathname.split('/')[2];
    switch (current) {
      case 'authorization':
        return SettingRoutes.authorization;
      default:
        return SettingRoutes.profile;
    }
  }, [location.pathname]);

  return (
    <main>
      <Grid container spacing={2}>
        <Grid md={3} xs={12}>
          <List>
            <NavigationItem
              icon={<PersonIcon />}
              label="用户信息"
              to="/settings"
              selected={currentRoute === SettingRoutes.profile}
            />
            <NavigationItem
              icon={<KeyIcon />}
              label="权限"
              to="/settings/authorization"
              selected={currentRoute === SettingRoutes.authorization}
            />
          </List>
        </Grid>
        <Grid md={9} xs={12}>
          <Box sx={{ pt: 2 }}>
            <Outlet />
          </Box>
        </Grid>
      </Grid>
    </main>
  );
}
