import useDialog from '@/common/use-dialog';
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Stack,
  TextField,
} from '@mui/material';
import { useFormik } from 'formik';
import { useEffect } from 'react';
import AddIcon from '@mui/icons-material/Add';
import * as Yup from 'yup';
import useFieldProperties from '@/routes/hooks/use-field-properties';
import { api } from '@/api/api';
import useToken from '@/common/user/use-token';
import useAppList from '../app-list/hooks/use-app-list';

export default function NewAppButton() {
  const { token } = useToken();
  const { mutate } = useAppList();
  const { isOpen, onClose, onOpen } = useDialog();
  const formik = useFormik({
    initialValues: {
      title: '',
      description: '',
    },
    validationSchema: Yup.object({
      title: Yup.string().required('App名称不能为空'),
      description: Yup.string(),
    }),
    async onSubmit(values) {
      await api.POST('/api/app', {
        body: {
          title: values.title,
          description: values.description,
          id: 0,
          hostnames: [],
        },
        params: {
          header: {
            authorization: token,
          },
        },
      });
      mutate();
      onClose();
    },
  });

  const { getFieldProperties } = useFieldProperties(formik);

  useEffect(() => {
    formik.resetForm();
  }, [isOpen]);

  return (
    <>
      <Button variant="contained" startIcon={<AddIcon />} onClick={onOpen}>
        NEW
      </Button>
      <Dialog
        open={isOpen}
        onClose={onClose}
        fullWidth
        maxWidth="sm"
        aria-labelledby="退出登录"
        aria-describedby="确认退出登录"
      >
        <form onSubmit={formik.handleSubmit}>
          <DialogTitle>新建 App</DialogTitle>
          <DialogContent>
            <Stack gap={2}>
              <TextField
                label="App 名称"
                variant="standard"
                fullWidth
                {...getFieldProperties('title')}
              />
              <TextField
                label="简介"
                variant="standard"
                fullWidth
                multiline
                minRows={3}
                maxRows={5}
                {...getFieldProperties('description')}
              />
            </Stack>
          </DialogContent>
          <DialogActions>
            <Button onClick={onClose}>取消</Button>
            <Button variant="contained" type="submit">
              提交
            </Button>
          </DialogActions>
        </form>
      </Dialog>
    </>
  );
}
