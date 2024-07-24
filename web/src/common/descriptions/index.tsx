import { Box, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2/Grid2';

import { ReactNode } from 'react';

interface Props {
  items: {
    label: ReactNode;
    value: ReactNode;
  }[];
}

export default function Descriptions({ items }: Props) {
  return (
    <Grid container spacing={2}>
      {items.map((item, index) => (
        <Grid key={index} xs={12} md={4}>
          <Box>
            <Typography fontWeight="bold" sx={{ mb: 1 }}>
              {item.label}
            </Typography>
            <Typography>{item.value}</Typography>
          </Box>
        </Grid>
      ))}
    </Grid>
  );
}
