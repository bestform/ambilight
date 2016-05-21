Gaming Ambilight
================

This project is set up to watch a folder and from every jpeg that is written to that folder, create an average light color and send it
to a configured Philips Hue.

The plan is to watch a screenshot folder and let an addon populate the folder with a new screenshot every x seconds.

An Addon has been written to do this job for WoW, but it turned out to be not suitable, because WoW pauses for a few milliseconds when taking a screenshot.
So instead, the current solution is to use a tool like "AutoScreenCap" to do the dirty work.
Just use ambilight to watch the folder the screenshot tool uses to store its screenshots.
Every screenshot will change the light of your lamp to the current average color in your game.

To do so, start the tool with the following flags:

--path=PATH_TO_YOUR_SCREENSHOT_FOLDER
--username=USERNAME_TO_YOUR_HUE_NETWORK
--ip=IP_OF_YOUR_HUE_NETWORK

Then just play the game and take a screenshot whenever you feel like it (or let an external tool do this for you)

If you want to delete every image after it has been processed, just add the option:
--delete

Vendored dependencies
---------------------

All dependencies that are not part of the standard library are vendored using govendor (https://github.com/kardianos/govendor)
They are already part of the repository, but to use them, you either need to use at least go1.6 or - if you are on go1.5 - please set the environment variable GO15VENDOREXPERIMENT to 1.
