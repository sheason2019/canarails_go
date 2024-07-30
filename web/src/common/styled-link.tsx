import {
  Link as RouterLink,
  LinkProps as RouterLinkProps,
} from '@modern-js/runtime/router';
import { Box, BoxProps } from '@mui/material';

type Props = BoxProps & RouterLinkProps;

export default function StyledLink(props: Props) {
  return (
    <Box
      component={RouterLink}
      {...props}
      sx={{
        color: 'inherit',
        textDecoration: 'none',
        userSelect: 'none',
      }}
    />
  );
}
