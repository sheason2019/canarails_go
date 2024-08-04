import { Link, Typography } from '@mui/material';
import { ReactNode } from 'react';

interface Props {
  label?: ReactNode;
  items: {
    label: string;
    link: string;
  }[];
}

export default function DependenciesGroup({ label, items }: Props) {
  return (
    <>
      {label && (
        <Typography sx={{ mt: 3, mb: 1 }} variant="h6">
          {label}
        </Typography>
      )}
      <Typography component="ul">
        {items.map(item => (
          <Typography key={item.label} sx={{ my: 1 }} component="li">
            <Link href={item.link} target="_blank" underline="none">
              {item.label}
            </Link>
          </Typography>
        ))}
      </Typography>
    </>
  );
}
