(function (global, document) {
  "use strict";

  (document);

  // Opens the moderatiosn settings window.
  global.moderationOpen = function() {
    $(".view").hide();
    $("#moderation-view").show();
    var data = global.dashboard.$data;
    var json = {
      "ownerToken": global.cookieGet("commentoOwnerToken"),
      "domain": data.domains[data.cd].domain
    }
    global.post(global.origin + "/api/comment/owner/list", json, function(resp) {
      if (!resp.success) {
        global.globalErrorShow(resp.message);
        return
      }
      for(var i in resp.comments) {
        resp.comments[i].creationDate = new Date(resp.comments[i].creationDate).toLocaleString()
      }
      Vue.set(data.domains[data.cd], "pending", resp.comments)
      data.domains[data.cd].pendingCommenters = resp.commenters
    });
    global.post(global.origin + "/api/comment/owner/listAll", json, function(resp) {
      if (!resp.success) {
        global.globalErrorShow(resp.message);
        return
      }
      for(var i in resp.comments) {
        resp.comments[i].creationDate = new Date(resp.comments[i].creationDate).toLocaleString()
      }
      Vue.set(data.domains[data.cd], "commentsAll", resp.comments)
      Vue.set(data.domains[data.cd], "commentersAll", resp.commenters)
    });
  };

  global.moderatorRemoveComment = function(hex) {
    var data = global.dashboard.$data;
    var indexToRemove = -1;
    for (var i in data.domains[data.cd].pending) {
      if (data.domains[data.cd].pending[i].commentHex === hex) {
        indexToRemove = i;
      }
    }
    data.domains[data.cd].pending.splice(indexToRemove, 1);
  }

  // Approves a comment
  global.moderatorApproveCommentHandler = function(hex) {
    var json = {
      "ownerToken": global.cookieGet("commentoOwnerToken"),
      "commentHex": hex
    }
    global.post(global.origin + "/api/comment/owner/approve", json, function(resp) {
      if (!resp.success) {
        global.globalErrorShow(resp.message);
        return
      }
      global.moderatorRemoveComment(hex)
    })
  }


  // Deletes a comment
  global.moderatorDeleteCommentHandler = function(hex) {
    var json = {
      "ownerToken": global.cookieGet("commentoOwnerToken"),
      "commentHex": hex
    }
    global.post(global.origin + "/api/comment/owner/delete", json, function(resp) {
      if (!resp.success) {
        global.globalErrorShow(resp.message);
        return
      }
      global.moderatorRemoveComment(hex)
    })
  }

  // Adds a moderator.
  global.moderatorNewHandler = function() {
    var data = global.dashboard.$data;
    var email = $("#new-mod").val();
    
    var json = {
      "ownerToken": global.cookieGet("commentoOwnerToken"),
      "domain": data.domains[data.cd].domain,
      "email": email,
    }

    var idx = -1;
    for (var i = 0; i < data.domains[data.cd].moderators.length; i++) {
      if (data.domains[data.cd].moderators[i].email === email) {
        idx = i;
        break;
      }
    }

    if (idx === -1) {
      data.domains[data.cd].moderators.push({"email": email, "timeAgo": "just now"});
      global.buttonDisable("#new-mod-button");
      global.post(global.origin + "/api/domain/moderator/new", json, function(resp) {
        global.buttonEnable("#new-mod-button");

        if (!resp.success) {
          global.globalErrorShow(resp.message);
          return
        }

        global.globalOKShow("Added a new moderator!");
        $("#new-mod").val("");
        $("#new-mod").focus();
      });
    } else {
      global.globalErrorShow("Already a moderator.");
    }
  }


  // Deletes a moderator.
  global.moderatorDeleteHandler = function(email) {
    var data = global.dashboard.$data;
    
    var json = {
      "ownerToken": global.cookieGet("commentoOwnerToken"),
      "domain": data.domains[data.cd].domain,
      "email": email,
    }

    var idx = -1;
    for (var i = 0; i < data.domains[data.cd].moderators.length; i++) {
      if (data.domains[data.cd].moderators[i].email === email) {
        idx = i;
        break;
      }
    }

    if (idx !== -1) {
      data.domains[data.cd].moderators.splice(idx, 1);
      global.post(global.origin + "/api/domain/moderator/delete", json, function(resp) {
        if (!resp.success) {
          global.globalErrorShow(resp.message);
          return
        }

        global.globalOKShow("Removed!");
      });
    }
  }


  // Adds a domain owner.
  global.domainOwnerAddHandler = function() {
    var data = global.dashboard.$data;
    var email = $("#new-domain-owner").val();
    var addAsModerator = $("#new-domain-owner-add-moderator").is(":checked");
    
    var json = {
      "ownerToken": global.cookieGet("commentoOwnerToken"),
      "domain": data.domains[data.cd].domain,
      "email": email,
    }

    var idx = -1;
    if (data.domains[data.cd].owners) {
      for (var i = 0; i < data.domains[data.cd].owners.length; i++) {
        if (data.domains[data.cd].owners[i].email === email) {
          idx = i;
          break;
        }
      }
    }

    if (idx === -1) {
      global.buttonDisable("#new-domain-owner-button");
      global.post(global.origin + "/api/domain/owner/add", json, function(resp) {
        global.buttonEnable("#new-domain-owner-button");

        if (!resp.success) {
          global.globalErrorShow(resp.message);
          return
        }

        global.globalOKShow("Added a new domain owner!");
        $("#new-domain-owner").val("");
        $("#new-domain-owner").focus();

        // Also add as moderator if checkbox is checked
        if (addAsModerator) {
          var modJson = {
            "ownerToken": global.cookieGet("commentoOwnerToken"),
            "domain": data.domains[data.cd].domain,
            "email": email,
          };
          global.post(global.origin + "/api/domain/moderator/new", modJson, function(modResp) {
            if (!modResp.success && modResp.message !== "Already a moderator.") {
              global.globalErrorShow("Failed to add a moderator: " + modResp.message);
            }

            global.domainRefresh();
          });
        }
      });
    } else {
      global.globalErrorShow("Already a domain owner.");
    }
  }


  // Removes a domain owner.
  global.domainOwnerRemoveHandler = function(ownerHex) {
    var data = global.dashboard.$data;
    
    var json = {
      "ownerToken": global.cookieGet("commentoOwnerToken"),
      "domain": data.domains[data.cd].domain,
      "removeOwnerHex": ownerHex,
    }

    global.post(global.origin + "/api/domain/owner/remove", json, function(resp) {
      if (!resp.success) {
        global.globalErrorShow(resp.message);
        return
      }

      // Refresh the domain list to get updated owners
      global.domainRefresh(function() {
        global.globalOKShow("Removed domain owner!");
      });
    });
  }

} (window.commento, document));
