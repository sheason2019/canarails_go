import { Stack, Typography } from '@mui/material';
import AddAppVariantButton from './components/add-app-variant-button';
import AppVariantsTable from './components/app-variants-table';

export default function AppVariants() {
  return (
    <main>
      <Stack sx={{ my: 2 }} alignItems="center" direction="row">
        <Typography variant="h5" sx={{ flex: 1 }}>
          流量泳道
        </Typography>
        <AddAppVariantButton />
      </Stack>
      <AppVariantsTable />
    </main>
  );
}
