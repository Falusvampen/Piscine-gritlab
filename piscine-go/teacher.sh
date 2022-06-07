interviewn=`grep -H "license" interviews/* | grep "\"" | cut -f1 -d ":" | rev | cut -f1 -d "-" | rev`
interview="cat interviews/interview-$interviewn"
export interviewnumber=$interviewn
echo $interviewnumber
$interview
echo $MAIN_SUSPECT