#!/bin/bash
/bin/api-server &
npm run dev &
wait -n
exit $?
