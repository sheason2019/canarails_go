import useDialog from '@/common/use-dialog';
import useFieldProperties from '@/routes/hooks/use-field-properties';
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  TextField,
} from '@mui/material';
import { useFormik } from 'formik';
import * as Yup from 'yup';
import useApp from '../../hooks/use-app';
import { useEffect } from 'react';
import useApi from '@/common/use-api';

export default function AddAppHostnameButton() {
  const api = useApi();
  const { data, mutate } = useApp();
  const { isOpen, onClose, onOpen } = useDialog();
  const formik = useFormik({
    initialValues: {
      hostname: '',
    },
    validationSchema: Yup.object({
      hostname: Yup.string().required('域名不能为空'),
    }),
    onSubmit: async values => {
      const app = data?.data;
      if (!app) throw new Error('app not exist');

      await api.PUT('/api/app/{id}', {
        params: {
          path: { id: app.id },
        },
        body: {
          ...app,
          hostnames: [...app.hostnames, values.hostname],
        },
      });
      onClose();
      mutate();
    },
  });

  const { getFieldProperties } = useFieldProperties(formik);

  useEffect(() => {
    formik.resetForm();
  }, [isOpen]);

  return (
    <>
      <Button variant="contained" onClick={onOpen}>
        新增
      </Button>
      <Dialog open={isOpen} onClose={onClose} fullWidth>
        <form onSubmit={formik.handleSubmit}>
          <DialogTitle>新增域名匹配</DialogTitle>
          <DialogContent>
            <TextField
              variant="standard"
              label="域名"
              fullWidth
              {...getFieldProperties('hostname')}
            />
          </DialogContent>
          <DialogActions>
            <Button variant="contained" type="submit">
              提交
            </Button>
          </DialogActions>
        </form>
      </Dialog>
    </>
  );
}
