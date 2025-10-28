


function renderCalendar(firstWeekday,daysInMonth) {

    week_day_administrator = firstWeekday + 1 

    const initialSpaces = 1 + firstWeekday*3

    console.log("日 月 火 水 木 金 土")
    process.stdout.write(" ".repeat(initialSpaces) + "1");

    if (week_day_administrator === 7) {
        console.log()
        week_day_administrator = 0;
    }

    for (i= 2 ; i <= daysInMonth; i++){
        // if (i >= 0 && i < 10) {
        //     if (week_day_administrator === 6) {
        //         week_day_administrator = 0
        //         // process.stdout.write(" ".repeat(2) + i);
        //         process.stdout.write(String(i).padStart(3));
        //         console.log()
        //     }
        //     else if(week_day_administrator === 0) {
        //         // process.stdout.write(" ".repeat(1) + i);
        //         process.stdout.write(String(i).padStart(2));
        //         week_day_administrator += 1
        //     }
        //     else{
        //         // process.stdout.write(" ".repeat(2) + i);
        //         process.stdout.write(String(i).padStart(3));
        //         week_day_administrator += 1
        //     }


        //     // else {
        //     //     // process.stdout.write(" ".repeat(1) + i);
        //     //     process.stdout.write(String(i).padStart(3," "));
        //     //     week_day_administrator += 1
        //     // }
  
        // }
        // else {

        //     if (week_day_administrator === 6) {
        //         week_day_administrator = 0
        //         // process.stdout.write(" ".repeat(1) + i);
        //         process.stdout.write(String(i).padStart(3));
        //         console.log()
        //     }
        //     else if(week_day_administrator === 0) {
        //         // process.stdout.write(" ".repeat() + i);
        //         process.stdout.write(String(i).padStart(2));
        //         week_day_administrator += 1
        //     }
        //     else{
        //         // process.stdout.write(" ".repeat(1) + i);
        //         process.stdout.write(String(i).padStart(3));
        //         week_day_administrator += 1

        //     }

        // }



        if (week_day_administrator === 6) {
            week_day_administrator = 0
            // process.stdout.write(" ".repeat(1) + i);
            process.stdout.write(String(i).padStart(3));
            console.log()
        }
        else if(week_day_administrator === 0) {
            // process.stdout.write(" ".repeat() + i);
            process.stdout.write(String(i).padStart(2));
            week_day_administrator += 1
        }
        else{
            // process.stdout.write(" ".repeat(1) + i);
            process.stdout.write(String(i).padStart(3));
            week_day_administrator += 1

        }

    }
//     for (let i = 1; i <= daysInMonth; i++) {
//   process.stdout.write(String(i).padStart(2, " ") + " ");

//   if (week_day_administrator === 6) {
//     console.log();
//     week_day_administrator = 0;
//   } else {
//     week_day_administrator++;
//   }
// }

    
}

const today = new Date(); // 今日の日付を取得
let lastDayOfMonth = new Date(today.getFullYear(), today.getMonth() + 1, 0); // 当月の翌月の0日を取得
let firstDayOfMonth = new Date(today.getFullYear(), today.getMonth() , 1);


// const args = process.argv;
// const command = args[2];
// const argMonth = args[3];

const [, ,  command, argMonth] = process.argv

// Case 1: `-m` が正しく指定されていて、月も1〜12
if (command === "-m") {
    const month = parseInt(argMonth);
    if (isNaN(month)) {
        console.log("-mの後に1から12の引数を指定してください");
    } else if (month < 1 || month > 12) {
        console.log("1〜12の範囲で月を指定してください。");
    } else {
        const currentYear = today.getFullYear();
        lastDayOfMonth = new Date(currentYear, month, 0);
        firstDayOfMonth = new Date(currentYear, month - 1, 1);

        console.log("     " + month + "月  " + currentYear);
        renderCalendar(firstDayOfMonth.getDay(), lastDayOfMonth.getDate());
    }

// Case 2: 引数がなかったら現在の月を表示
} else if (!command) {
    const currentYear = today.getFullYear();
    const currentMonth = today.getMonth() + 1;

    console.log("     " + currentMonth + "月  " + currentYear);
    renderCalendar(firstDayOfMonth.getDay(), lastDayOfMonth.getDate());

// Case 3: `-m` 以外が指定された場合
} else {
    console.log("正しい引数 -m を指定してください");
}
