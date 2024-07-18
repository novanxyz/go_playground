const newman = require('newman'); // require newman in your project

// call newman.run to pass `options` object and wait for callback
newman.run({
    collection: [
        require('./scenarios/create-a-new-task.json'),
        require('./scenarios/delete-a-specific-task-by-id.json')
            ], 
    reporters: 'cli',
    environment: require('./envvars.json')
}, function (err) {
	if (err) { throw err; }
    console.log('collection run complete!');
});