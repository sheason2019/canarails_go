import {
  Link as RouterLink,
  LinkProps as RouterLinkProps,
} from '@modern-js/runtime/router';
import { Link, LinkProps } from '@mui/material';

type Props = LinkProps & RouterLinkProps;

export default function StyledLink(props: Props) {
  return <Link {...props} component={RouterLink} />;
}
