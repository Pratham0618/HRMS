import  rest from 'msw';

export const handlers = [
  rest.getResponse('https://my-json-server.typicode.com/pk2601/employee-dashboard/posts', (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json([
        {
          key: '1',
          employeeid: 1,
          name: 'John Brown',
          departmentid: 101,
          managerid: 201,
          email: 'john.brown@example.com',
          phone: '123-456-7890',
          address: 'New York No. 1 Lake Park',
          dob: '1990-01-01',
        },
        {
          key: '2',
          employeeid: 2,
          name: 'Jim Green',
          departmentid: 102,
          managerid: 202,
          email: 'jim.green@example.com',
          phone: '098-765-4321',
          address: 'London No. 1 Lake Park',
          dob: '1982-05-12',
        },
        {
          key: '3',
          employeeid: 3,
          name: 'Joe Black',
          departmentid: 103,
          managerid: 203,
          email: 'joe.black@example.com',
          phone: '111-222-3333',
          address: 'Sydney No. 1 Lake Park',
          dob: '1990-03-15',
        },
        {
          key: '4',
          employeeid: 4,
          name: 'Jane Doe',
          departmentid: 104,
          managerid: 204,
          email: 'jane.doe@example.com',
          phone: '444-555-6666',
          address: 'San Francisco No. 2 Lake Park',
          dob: '1993-09-21',
        },
      ])
    );
  }),

  // rest.post('https://my-json-server.typicode.com/pk2601/employee-dashboard/posts', (req, res, ctx) => {
  //   const { name, departmentid, managerid, email, phone, address, dob } = req.body;
  //   const newEmployee = {
  //     key: String(Math.random()),
  //     employeeid: Math.floor(Math.random() * 1000) + 1,
  //     name,
  //     departmentid,
  //     managerid,
  //     email,
  //     phone,
  //     address,
  //     dob,
  //   };
  //   return res(
  //     ctx.status(201),
  //     ctx.json(newEmployee)
  //   );
  // }),
];

